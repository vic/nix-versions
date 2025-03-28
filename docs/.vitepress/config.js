export default {
  themeConfig: {
    sidebar: {
      // This sidebar gets displayed when a user
      // is on `guide` directory.
      '/guide/': [
        {
          text: 'Guide',
          items: [
            { text: 'Installing', link: '/guide/installing' },
            { text: 'Listing Versions', link: '/guide/listing-versions' },
            { text: 'Finding Packages', link: '/guide/finding-packages' },
            { text: 'Output formats', link: '/guide/finding-packages' },
            { text: 'Command line options', link: '/guide/finding-packages' }
          ]
        },
        {text: 'Integrations', link: '/integrations/nix'},
        {text: 'Contributing', link: '/contributing/why'},
      ],

      '/integrations/': [
        {text: 'Guide', link: '/guide/installing'},
        {
          text: 'Integrations',
          items: [
              {
                  text: 'Standalone - Easiest Dev Env',
                  items: [
                    { text: 'nix installables', link: '/integrations/nix' },
                    { text: 'direnv', link: '/integrations/direnv' },
                  ]
              },
              {
                  text: 'With other nix configurations',
                  items: [
                    { text: 'fetchPackageAtVersion', link: '/integrations/fetcher' },
                    { text: 'flakes', link: '/integrations/flakes' },
                    { text: 'flake-parts', link: '/integrations/flake-parts' },
                  ]
              },
              {
                  text: 'With advanced nix environments',
                  items: [
                    { text: 'devenv', link: '/integrations/devenv' },
                    { text: 'devshell', link: '/integrations/devshell'},
                    { text: 'NixOS', link: '/integrations/devshell'},
                    { text: 'home-manager', link: '/integrations/devshell'},
                    { text: 'nix-darwin', link: '/integrations/devshell'},
                  ]
              }
          ]
        }
      ]
    }
  }
}
