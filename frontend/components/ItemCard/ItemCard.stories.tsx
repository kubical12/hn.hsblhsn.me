import { ComponentStory } from '@storybook/react'
import { ItemCard } from './ItemCard'

export default {
  title: 'ItemCard',
  component: ItemCard,
}

const Template: ComponentStory<typeof ItemCard> = (args) => (
  <ItemCard {...args} />
)

export const TopStory = Template.bind({})
TopStory.args = {
  item: {
    id: '30959025',
    title: 'NotepadNext: A cross-platform reimplementation of Notepad++',
    url: 'https://github.com/dail8859/NotepadNext',
    descendants: 82,
    score: 117,
    text: '',
    openGraph: {
      description:
        'A cross-platform, reimplementation of Notepad++. Contribute to dail8859/NotepadNext development by creating an account on GitHub.',
      image: [
        {
          url: 'https://repository-images.githubusercontent.com/224468265/85a97400-96c9-11eb-9836-e8abb8e901b9',
          alt: '',
          width: 0,
        },
      ],
    },
    __typename: 'Story',
  },
}

export const AskStory = Template.bind({})
AskStory.args = {
  item: {
    id: '30961400',
    title: 'Ask HN: Is there a way to get Agile right?',
    text: 'My team has recently began considering an agile methodology of tracking work efforts, and while we are not strictly a development team, it appears way more promising than managing work efforts through a shared spread sheet.<p>However, any attempt to research ways to adopt this sort of framework leads me to articles where people are expressing how horribly wrong it can go.<p>So, my question: Is there a way to get Agile right? or are we doomed to fail in the same ways that has been described by countless articles discussing the downfalls of Agile?',
    url: 'https://github.com/dail8859/NotepadNext',
    descendants: 5,
    score: 12,
    openGraph: {},
    __typename: 'Story',
  },
}
