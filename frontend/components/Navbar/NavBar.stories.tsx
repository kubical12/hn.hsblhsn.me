import { ComponentStory } from '@storybook/react'
import { NavBar } from './NavBar'

export default {
  title: 'NavBar',
  component: NavBar,
}

const Template: ComponentStory<typeof NavBar> = (args) => <NavBar {...args} />

export const Default = Template.bind({})
