import { ComponentStory } from '@storybook/react'
import { ItemCardList } from './ItemCardList'

export default {
  title: 'ItemCardList',
  component: ItemCardList,
}

const Template: ComponentStory<typeof ItemCardList> = (args) => (
  <ItemCardList {...args} />
)

export const Default = Template.bind({})
Default.args = {
  items: {
    edges: [
      {
        node: {
          __typename: 'Story',
          id: '30977164',
          type: 'story',
          score: 27,
          descendants: 1,
          openGraph: {
            title:
              'GitHub - nikitinprior/dF80: Restored F80 compiler code for CP/M',
            description:
              'Restored F80 compiler code for CP/M. Contribute to nikitinprior/dF80 development by creating an account on GitHub.',
            image: [
              {
                url: 'https://opengraph.githubassets.com/0091aaa8d56f3f7f96f7f6c36dfe341c37511e11d1f37b32084a8dca46ff6622/nikitinprior/dF80',
                alt: '',
                width: 0,
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Story',
          id: '30975378',
          type: 'story',
          score: 195,
          descendants: 181,
          openGraph: {
            title: 'Do Cats Have Intelligence/How Intelligent Are Cats?',
            description: '',
            image: [],
          },
        },
      },
      {
        node: {
          __typename: 'Story',
          id: '30975532',
          type: 'story',
          score: 141,
          descendants: 65,
          openGraph: {
            title: 'W3M Rocks',
            description:
              'The finer aspects of W3M, the terminal web browser and pager',
            image: [],
          },
        },
      },
      {
        node: {
          __typename: 'Story',
          id: '30955023',
          type: 'story',
          score: 43,
          descendants: 23,
          openGraph: {
            title: 'Blog - How Swipe Typing works',
            description:
              'The input methods supported by software keyboards are "tap typing" and "swipe typing" where users swipe their finger across the keyboard.',
            image: [
              {
                url: 'https://www.fleksy.com/wp-content/uploads/2022/04/how-swipe-typing-works.png',
                alt: '',
                width: 0,
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Story',
          id: '30976337',
          type: 'story',
          score: 73,
          descendants: 22,
          openGraph: {
            title: '∑ Xah Code',
            description: '',
            image: [],
          },
        },
      },
      {
        node: {
          __typename: 'Story',
          id: '30965572',
          type: 'story',
          score: 26,
          descendants: 38,
          openGraph: {
            title: 'Why Compilers Don’t Autocorrect “Obvious” Parse Errors',
            description:
              'Last month, someone on Twitter relayed a conversation with their 8 year old daughter, who is learning Python. The kid wants to know “If the computer knows I’m missing a semicolon here, …',
            image: [
              {
                url: 'https://chelseatroy.com/wp-content/uploads/2022/03/Screen-Shot-2022-03-11-at-9.04.11-AM.png',
                alt: '',
                width: 0,
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Story',
          id: '30977147',
          type: 'story',
          score: 275,
          descendants: 417,
          openGraph: {
            title: 'Heresy',
            description: '',
            image: [],
          },
        },
      },
      {
        node: {
          __typename: 'Story',
          id: '30966362',
          type: 'story',
          score: 54,
          descendants: 11,
          openGraph: {
            title:
              'GitHub - zineland/zine: Zine - a simple and opinionated tool to build your own magazine.',
            description:
              'Zine - a simple and opinionated tool to build your own magazine. - GitHub - zineland/zine: Zine - a simple and opinionated tool to build your own magazine.',
            image: [
              {
                url: 'https://repository-images.githubusercontent.com/463489443/841f79ec-570b-4a52-b855-2ca4852acd11',
                alt: '',
                width: 0,
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Story',
          id: '30965343',
          type: 'story',
          score: 69,
          descendants: 11,
          openGraph: {
            title: 'My Org Roam Notes Workflow - Hugo Cisneros',
            description:
              'A description of my org-roam notes workflow, from writing the notes to publishing them on my website.',
            image: [
              {
                url: '/img/note-graph.jpg',
                alt: '',
                width: 0,
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Story',
          id: '30976616',
          type: 'story',
          score: 79,
          descendants: 32,
          openGraph: {
            title: 'Erik Naggum - Wikiquote',
            description: '',
            image: [
              {
                url: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/69/ErikNaggum_LUGM99_DKL.jpg/1200px-ErikNaggum_LUGM99_DKL.jpg',
                alt: '',
                width: 0,
              },
              {
                url: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/69/ErikNaggum_LUGM99_DKL.jpg/800px-ErikNaggum_LUGM99_DKL.jpg',
                alt: '',
                width: 0,
              },
              {
                url: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/69/ErikNaggum_LUGM99_DKL.jpg/640px-ErikNaggum_LUGM99_DKL.jpg',
                alt: '',
                width: 0,
              },
            ],
          },
        },
      },
    ],
  },
}
