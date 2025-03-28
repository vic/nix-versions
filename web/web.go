package main

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/vic/ntv/packages/flake"
	"github.com/vic/ntv/packages/search"
	"github.com/vic/ntv/packages/search_spec"
)

func Web() {
	// http.HandleFunc("/default.nix/", HandleDefaultNix)
	http.HandleFunc("/flake.nix/", HandleFlakeNix)
	http.HandleFunc("/flake.zip/", HandleFlakeZip)
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

func createFlake(args []string) (*flake.Context, error) {
	specs, err := search_spec.ParseSearchSpecs(args, nil)
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

	fmt.Printf("Checking Nixpkgs: %s\n", url)
	h := sha256.New()
	_, err = io.Copy(h, resp.Body)
	if err != nil {
		return "", err
	}
	sum := h.Sum(nil)
	base64Url := base64.URLEncoding.EncodeToString(sum)
	sri := "sha256-" + base64Url
	fmt.Printf("Nixpkgs: %s\nSRI: %s\n", url, sri)
	return sri, nil
}
