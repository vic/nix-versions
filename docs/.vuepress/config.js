//import { defaultTheme } from '@vuepress/theme-default'
import { hopeTheme} from "vuepress-theme-hope";
import { defineUserConfig } from 'vuepress'
import { webpackBundler } from '@vuepress/bundler-webpack'

export default defineUserConfig({
  lang: 'en-US',

  title: '> nix-versions',
  description: 'Any version of <code>200_000+</code> nix packages at your fingerprints.',

  theme: hopeTheme({
    // logo: 'https://github.com/user-attachments/assets/307fd2a5-7bf5-45bc-8436-ce6fec1f914a',
    navbar: ['/', '/guide/installing.html'],
  }),

  bundler: webpackBundler(),
})
