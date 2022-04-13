import { ComponentStory } from '@storybook/react'
import { Comment } from './Comment'

export default {
  title: 'Comment',
  component: Comment,
}

const Template: ComponentStory<typeof Comment> = (args) => <Comment {...args} />

export const Default = Template.bind({})
Default.args = {
  comment: {
    id: '30978624',
    text: 'Probably it&#x27;s just me, but after 2016 and all that idiotic internet edgelord-ery that enabled and accompanied it, and the repercussions we will still feel for decades, I&#x27;ve lost all patience and interest with these type of 14-and-this-is-deep edgelords, especially when they&#x27;re in their 20s or even older.<p>They don&#x27;t add anything that for instance Houellebecq hasn&#x27;t formulated before them, and way more poignantly. Just because they don&#x27;t read literature but stew in the muddy waters of forums and usenet and other places, doesn&#x27;t mean their pontificating can be excused.',
    by: 'suction',
    time: 1649607925,
    parent: '30976616',
    __typename: 'Comment',
    descendants: 120,
  },
}
