// @ts-check
import { defineConfig, fontProviders } from "astro/config";
import starlight from "@astrojs/starlight";

import mermaid from "astro-mermaid";
import catppuccin from "@catppuccin/starlight";

// https://astro.build/config
export default defineConfig({
  experimental: {
    fonts: [
      {
        provider: fontProviders.google(),
        name: "Victor Mono",
        cssVariable: "--sl-font",
      },
    ],
  },
  integrations: [
    mermaid({
      theme: "forest",
      autoTheme: true,
    }),
    starlight({
      title: "nix-version",
      sidebar: [
        {
          label: "Getting Started",
          items: [
            { label: "Installing", slug: "getting-started/installing" },
            { label: "Listing package versions", slug: "getting-started/listing-versions" },
            { label: "Finding packages", slug: "getting-started/finding-packages" },
            { label: "CLI Help", slug: "getting-started/cli-help" },
          ],
        },
        {
          label: "Guides",
          items: [
            { label: "Tools Version Manager", slug: "tools-version-manager" },
            { label: "Flake Generator", slug: "flake-generator" },
            { label: "Non-Flake Environments", slug: "non-flake-environments" },
            { label: "Installing packages on profiles", slug: "installing-packages-on-profiles" },
          ],
        },
      ],
      components: {
        Sidebar: "./src/components/Sidebar.astro",
        Footer: "./src/components/Footer.astro",
        SocialIcons: "./src/components/SocialIcons.astro",
        PageSidebar: "./src/components/PageSidebar.astro",
      },
      plugins: [
        catppuccin({
          dark: { flavor: "macchiato", accent: "mauve" },
          light: { flavor: "latte", accent: "mauve" },
        }),
      ],
      editLink: {
        baseUrl: "https://github.com/vic/flake-file/edit/main/docs/",
      },
      customCss: ["./src/styles/custom.css"],
    }),
  ],
});
