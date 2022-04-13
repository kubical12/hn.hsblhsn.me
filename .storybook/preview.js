import {addDecorator} from '@storybook/react';
import withProviders from './prelude';
addDecorator(withProviders);
export const parameters = {
  actions: { argTypesRegex: "^on[A-Z].*" },
  controls: {
    matchers: {
      color: /(background|color)$/i,
      date: /Date$/,
    },
  },
}