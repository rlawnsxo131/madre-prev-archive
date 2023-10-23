import '../src/styles/global.css';

import React from 'react';

import type { Preview } from '@storybook/react';
import { addons } from '@storybook/preview-api';

import { DARK_MODE_EVENT_NAME } from 'storybook-dark-mode';

const channel = addons.getChannel();

const preview: Preview = {
  parameters: {
    actions: { argTypesRegex: '^on[A-Z].*' },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/,
      },
    },
  },
  decorators: [
    (Story) => {
      React.useEffect(() => {
        const handler = (isDarkMode: boolean) => {
          document
            .querySelector(':root')
            ?.setAttribute('data-theme', isDarkMode ? 'dark' : 'light');
        };

        channel.on(DARK_MODE_EVENT_NAME, handler);
        return () => channel.off(DARK_MODE_EVENT_NAME, handler);
      }, [channel]);

      return <Story />;
    },
  ],
};

export default preview;
