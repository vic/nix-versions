package main

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/vic/ntv/packages/flake"
	"github.com/vic/ntv/packages/search"
	"github.com/vic/ntv/packages/search_spec"
)

func main() {
	http.HandleFunc("/flake.nix/", HandleFlakeNix)
	http.HandleFunc("/flake.zip/", HandleFlakeZip)
	http.HandleFunc("/default.nix/", HandleDefaultNix)
	http.HandleFunc("/default.zip/", HandleDefaultZip)
	http.HandleFunc("/nixpkgs-sri/", HandleSri)
	fs := http.FileServer(http.Dir("www/"))
	http.Handle("/", http.StripPrefix("/", fs))
	addr := os.ExpandEnv(":$PORT")
	if addr == ":" {
		addr = ":8080"
	}
	fmt.Println("Listening on", addr)
	http.ListenAndServe(addr, nil)
}

func searchSpecs(args []string) (search.PackageSearchResults, error) {
	backend := search_spec.VersionsBackend{
		NixHub: &search_spec.Unit{},
	}
	specs, err := search_spec.ParseSearchSpecs(args, backend)
	if err != nil {
		return nil, err
	}

	res, err := search.PackageSearchSpecs(specs).Search()
	if err != nil {
		return nil, err
	}

	if err := res.EnsureOneSelected(); err != nil {
		return nil, err
	}
	if err := res.EnsureUniquePackageNames(); err != nil {
		return nil, err
	}
	return res, nil
}

func createFlake(args []string) (*flake.Context, error) {
	res, err := searchSpecs(args)
	if err != nil {
		return nil, err
	}

	f := flake.New()
	for _, r := range res {
		f.AddTool(r)
	}

	return f, nil
}

func renderFlake(args []string) (string, error) {
	f, err := createFlake(args)
	if err != nil {
		return "", err
	}

	code, err := f.Render(false)
	if err != nil {
		return "", err
	}

	return code, nil
}

type Prefetch struct {
	tool *search.PackageSearchResult
	Url  string
	Hash string
}

func prefetch(res search.PackageSearchResults) ([]Prefetch, error) {
	group, _ := errgroup.WithContext(context.Background())
	results := make([]Prefetch, len(res))
	for i, s := range res {
		i, s := i, s
		group.Go(func() error {
			// url := "https://codeload.github.com/nixos/nixpkgs/tar.gz/refs/" + s.Selected.Revision
			url := "https://github.com/NixOS/nixpkgs/archive/" + s.Selected.Revision + ".tar.gz"
			hash, err := Sri(url)
			if err != nil {
				return err
			}
			results[i] = Prefetch{
				Url:  url,
				Hash: hash,
				tool: s,
			}
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func unpackArray[S ~[]E, E any](s S) []any {
	r := make([]any, len(s))
	for i, e := range s {
		r[i] = e
	}
	return r
}
func writeIndent(buff *bytes.Buffer) func(int, string, ...string) {
	return func(i int, s string, x ...string) {
		if len(x) == 0 {
			buff.WriteString((strings.Repeat("  ", i)) + s + "\n")
		} else {
			buff.WriteString((strings.Repeat("  ", i)) + fmt.Sprintf(s, unpackArray(x)...) + "\n")
		}
	}
}

func renderDefaultNix(args []string) (string, error) {
	res, err := searchSpecs(args)
	if err != nil {
		return "", err
	}

	fetched, err := prefetch(res)
	if err != nil {
		return "", err
	}

	buff := bytes.Buffer{}
	w := writeIndent(&buff)
	w(0, "{ pkgs ? import <nixpkgs> {} }:")
	w(0, "let")
	for _, p := range fetched {
		w(1, "tools.\"%s\" = {", p.tool.Selected.Name)
		w(2, "url = \"%s\";", p.Url)
		w(2, "hash = \"%s\";", p.Hash)
		w(2, "version = \"%s\";", p.tool.Selected.Version)
		w(2, "attrPath = \"%s\";", p.tool.Selected.Attribute)
		w(1, "};")
	}
	outs := `
	nixpkgs = pkgs.lib.mapAttrs (name: tool:
	  let 
	    archive = pkgs.fetchurl { 
		  name = "${name}-${tool.version}-nixpkgs.tar.gz";
	      url = tool.url;
	      hash = tool.hash;
	    };
		# since our sri was computed on the tarball, we need to unpack it
		# we could use fetchzip, but we'd have to extract the tarball on 
		# the server to compute recursive sri (which we wont do).
		unpacked = pkgs.stdenvNoCC.mkDerivation {
		  name = "${name}-${tool.version}-nixpkgs";
		  phases = [ "unpackPhase" ];
		  unpackPhase = ''
		    unpackFile ${archive}
			mv nixpkgs-* $out
		  '';
		};
	  in unpacked
	) tools;
	packages = pkgs.lib.mapAttrs (name: tool:
	  let 
	    pkgs' = import nixpkgs.${name} { inherit (pkgs) system config; };
	    path = pkgs.lib.splitString "." tool.attrPath;
	    pkg = pkgs.lib.getAttrFromPath path pkgs';
	  in pkg
	) tools;
	pkgsEnv = pkgs.buildEnv { name = "tools"; paths = pkgs.lib.attrValues packages; }; 
	devShell = pkgs.mkShell { buildInputs = [ pkgsEnv ]; };
    `
	w(0, strings.ReplaceAll(outs, "\t", "  "))
	w(0, "in { inherit tools nixpkgs packages pkgsEnv devShell; }")

	return buff.String(), nil
}

func HandleDefaultNix(w http.ResponseWriter, r *http.Request) {
	werr := func(err error) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	path := strings.TrimPrefix(r.URL.Path, "/default.nix/")
	parts := strings.Split(path, "/")
	fmt.Println("Gen default.nix: ", parts)

	nix, err := renderDefaultNix(parts)
	if err != nil {
		werr(err)
		return
	}

	w.Header().Set("Content-Type", "text/x-nix")
	w.Header().Set("Content-Disposition", "attachment; filename=default.nix")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	fmt.Fprint(w, nix)
}

func HandleFlakeNix(w http.ResponseWriter, r *http.Request) {
	werr := func(err error) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	path := strings.TrimPrefix(r.URL.Path, "/flake.nix/")
	parts := strings.Split(path, "/")
	fmt.Println("Gen flake.nix: ", parts)

	flake, err := renderFlake(parts)
	if err != nil {
		werr(err)
		return
	}

	w.Header().Set("Content-Type", "text/x-nix")
	w.Header().Set("Content-Disposition", "attachment; filename=flake.nix")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	fmt.Fprint(w, flake)
}

func HandleDefaultZip(w http.ResponseWriter, r *http.Request) {
	werr := func(err error) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	path := strings.TrimPrefix(r.URL.Path, "/default.zip/")
	parts := strings.Split(path, "/")
	fmt.Println("Gen default.zip: ", parts)

	nix, err := renderDefaultNix(parts)
	if err != nil {
		werr(err)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=default.zip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	zw := zip.NewWriter(w)
	defer zw.Close()
	err = writeFileToZip(zw, "default.nix", []byte(nix))
	if err != nil {
		werr(err)
		return
	}
}

func HandleFlakeZip(w http.ResponseWriter, r *http.Request) {
	werr := func(err error) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	path := strings.TrimPrefix(r.URL.Path, "/flake.zip/")
	parts := strings.Split(path, "/")
	fmt.Println("Gen flake.zip: ", parts)

	flake, err := renderFlake(parts)
	if err != nil {
		werr(err)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=flake.zip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	zw := zip.NewWriter(w)
	defer zw.Close()
	err = writeFileToZip(zw, "flake.nix", []byte(flake))
	if err != nil {
		werr(err)
		return
	}
}

func writeFileToZip(zw *zip.Writer, path string, content []byte) error {
	w, err := zw.Create(path)
	if err != nil {
		return err
	}
	_, err = w.Write(content)
	return err
}

func HandleSri(w http.ResponseWriter, r *http.Request) {
	rev := strings.TrimPrefix(r.URL.Path, "/nixpkgs-sri/")
	sri, err := Sri("https://codeload.github.com/nixos/nixpkgs/zip/refs/heads/" + rev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, sri)
}

func Sri(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	h := sha256.New()
	_, err = io.Copy(h, resp.Body)
	if err != nil {
		return "", err
	}
	sum := h.Sum(nil)
	base64Url := base64.URLEncoding.EncodeToString(sum)
	sri := "sha256-" + strings.ReplaceAll(base64Url, "-", "+")
	return sri, nil
}

func Sha256(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	h := sha256.New()
	_, err = io.Copy(h, resp.Body)
	if err != nil {
		return "", err
	}
	sum := h.Sum(nil)
	return hex.EncodeToString(sum), nil
}
