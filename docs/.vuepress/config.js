import { hopeTheme} from "vuepress-theme-hope";
import { defineUserConfig } from 'vuepress'
import { viteBundler } from '@vuepress/bundler-vite'

export default defineUserConfig({
  lang: 'en-US',

  title: '> nix-versions',
  description: 'Any version of <code>200_000+</code> nix packages at your fingerprints.',

  theme: hopeTheme({
    // logo: 'https://github.com/user-attachments/assets/307fd2a5-7bf5-45bc-8436-ce6fec1f914a',
    navbar: ['/', '/getting-started/listing-versions.html', '/tools-version-manager.html', '/flake-generator.html'],
    outline: [1, 2, 3],
    pagePatterns: ["**/*.md", "!**/*.snippet.md", "!**/*.ansi.html", "!.vuepress", "!node_modules"],
    markdown: {
      include: true,
    },
  }),

  bundler: viteBundler(),
})
