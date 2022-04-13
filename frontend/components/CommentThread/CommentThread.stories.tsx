import { ComponentStory } from '@storybook/react'
import { CommentThread } from './CommentThread'

export default {
  title: 'CommentThread',
  component: CommentThread,
}

const Template: ComponentStory<typeof CommentThread> = (args) => (
  <CommentThread {...args} />
)

export const Default = Template.bind({})
Default.args = {
  comments: {
    edges: [
      {
        node: {
          __typename: 'Comment',
          id: '30978624',
          text: 'Probably it&#x27;s just me, but after 2016 and all that idiotic internet edgelord-ery that enabled and accompanied it, and the repercussions we will still feel for decades, I&#x27;ve lost all patience and interest with these type of 14-and-this-is-deep edgelords, especially when they&#x27;re in their 20s or even older.<p>They don&#x27;t add anything that for instance Houellebecq hasn&#x27;t formulated before them, and way more poignantly. Just because they don&#x27;t read literature but stew in the muddy waters of forums and usenet and other places, doesn&#x27;t mean their pontificating can be excused.',
          by: 'suction',
          time: 1649607925,
          parent: 30976616,
          comments: {
            edges: [],
          },
        },
      },
      {
        node: {
          __typename: 'Comment',
          id: '30976867',
          text: 'Apparently it&#x27;s smart troll day on HN. (Xah is just down the page right now.) It&#x27;s like a celebration of the kind of person I admired before I realized that other people are real too.',
          by: 'JasonFruit',
          time: 1649598258,
          parent: 30976616,
          comments: {
            edges: [
              {
                node: {
                  __typename: 'Comment',
                  id: '30977223',
                  text: 'Probably saw this guy mentioned on that awfully long Xah Lee wikipost.',
                  by: 'sva_',
                  time: 1649601278,
                },
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Comment',
          id: '30977181',
          text: 'I read comp.lang.lisp from 1998 - 2006. Lot&#x27;s of interesting technical material. Many of Erik&#x27;s technical posts are insightful. A few of these quotes make him look like a doofus keyboard warrior, but I could tell after reading his writing for so long that he was a genuine good guy, albeit at times a bit abrasive.<p>Can&#x27;t believe that it has been twenty years since I was introduced to common lisp. Think that I will go back to it after I retire.',
          by: 'vdas',
          time: 1649600951,
          parent: 30976616,
          comments: {
            edges: [
              {
                node: {
                  __typename: 'Comment',
                  id: '30978216',
                  text: 'At least Naggum&#x27;s was creative and entertaining in his abrasiveness. I&#x27;m not sure if I would have enjoyed actually interacting with him, but I always chuckle a bit reading through his old rants. If you read these things with a bit of humour it&#x27;s quite funny.',
                  by: 'Beltalowda',
                  time: 1649605929,
                },
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Comment',
          id: '30976883',
          text: '&gt; People search for the meaning of life, but this is the easy question: we are born into a world that presents us with many millennia of collected knowledge and information, and all our predecessors ask of us is that we not waste our brief life ignoring the past only to rediscover or reinvent its lessons badly<p>Yeah the Not Invented Here mindset is hard to avoid. So many people ignorant of the whole body of knowledge on various subjects before them, and the best they can do is re-invent the wheel and cry &#x27;Eureka!&#x27; (e.g Like when stumbling upon a yoga asana&#x2F;posture that was committed to text 2000 years ago or accidentally coding Bubble Sort with no prior knowledge of algorithms).',
          by: 'legrande',
          time: 1649598432,
          parent: 30976616,
          comments: {
            edges: [
              {
                node: {
                  __typename: 'Comment',
                  id: '30978383',
                  text: '&gt; accidentally coding Bubble Sort with no prior knowledge of algorithms<p>Heh, I did that when I was a kid.  I was so proud of myself, right up until I learned how bad it was.<p>If only there were some way for the tech world at large to learn that sort of lesson.',
                  by: 'rauhl',
                  time: 1649606671,
                },
              },
              {
                node: {
                  __typename: 'Comment',
                  id: '30977234',
                  text: 'We need better, more intelligent tools to search through the corpus of human knowledge.',
                  by: 'sva_',
                  time: 1649601390,
                },
              },
              {
                node: {
                  __typename: 'Comment',
                  id: '30976904',
                  text: 'Sometimes, re-inventing something is easier than trying to find it somewhere else.',
                  by: 'Hermel',
                  time: 1649598705,
                },
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Comment',
          id: '30976864',
          text: 'Erik died in 2009.<p>Erik Naggum in memoriam  <a href="https:&#x2F;&#x2F;perpelle.net&#x2F;artikler-og-leserinnlegg&#x2F;erik-naggum-in-memoriam&#x2F;" rel="nofollow">https:&#x2F;&#x2F;perpelle.net&#x2F;artikler-og-leserinnlegg&#x2F;erik-naggum-in...</a>',
          by: 'jzebedee',
          time: 1649598234,
          parent: 30976616,
          comments: {
            edges: [],
          },
        },
      },
      {
        node: {
          __typename: 'Comment',
          id: '30977214',
          text: 'His famous anti Perl rant circa 2000: (I worked as a Perl dev in 2000)\n<a href="https:&#x2F;&#x2F;groups.google.com&#x2F;g&#x2F;comp.lang.lisp&#x2F;c&#x2F;LGeQBt_ClfI&#x2F;m&#x2F;Y_iyHKvrdvwJ" rel="nofollow">https:&#x2F;&#x2F;groups.google.com&#x2F;g&#x2F;comp.lang.lisp&#x2F;c&#x2F;LGeQBt_ClfI&#x2F;m&#x2F;Y...</a>',
          by: 'mmaunder',
          time: 1649601219,
          parent: 30976616,
          comments: {
            edges: [
              {
                node: {
                  __typename: 'Comment',
                  id: '30977496',
                  text: 'This guy would have had a field day with python, if he were alive right now',
                  by: 'isoprophlex',
                  time: 1649602724,
                },
              },
              {
                node: {
                  __typename: 'Comment',
                  id: '30978471',
                  text: 'It&#x27;s... not very good? The bar for language takedowns is apparently much higher today than in 2000.',
                  by: 'tptacek',
                  time: 1649607051,
                },
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Comment',
          id: '30977028',
          text: 'I noticed on his websites that he likes to write dates in a format like &quot;2009-121&quot; or &quot;2006-257&quot;, i.e., year and then day of year (1..366).<p>I also use that concept, though I write it as XXYYY where XX is years post 2000 and YYY is the day number. For example, today is 22100.<p>Why do I like this? I find this is more &quot;mentally ergonomic&quot; to use for software tooling I have made for myself only. I usually don&#x27;t care about the month or date of the month; I care about offsets (e.g. tomorrow is 22101 and ten days from now is 22110). If I&#x27;m computing an offset in my head, I don&#x27;t have to account for wrap-arounds due to month (e.g. 10 days from March 25 is April 4 because March has 31 days.) The XXYYY format is about as concise as you can get (i.e., takes up minimal space on the screen) while still capturing all the information I care about. XXYYY sorts easily (though that&#x27;s also true of YY.MM.DD which is my second favorite format).',
          by: 'javert',
          time: 1649599841,
          parent: 30976616,
          comments: {
            edges: [
              {
                node: {
                  __typename: 'Comment',
                  id: '30978811',
                  text: 'What do you think about week-counting dates? They are quite popular for project planning in nordics.<p>You sometimes see them written like 22w10 or even 22w10.5  (2022, week 10, day 5, which is Friday).<p>Outlook etc support showing week numbering if you opt in.',
                  by: 'kzrdude',
                  time: 1649608921,
                },
              },
              {
                node: {
                  __typename: 'Comment',
                  id: '30977548',
                  text: 'This is actually specified in ISO 8601; they call it the ordinal date.',
                  by: 'rdpintqogeogsaa',
                  time: 1649602974,
                },
              },
              {
                node: {
                  __typename: 'Comment',
                  id: '30977466',
                  text: 'Have you considered that your system might not be communicative to others?',
                  by: 'JasonFruit',
                  time: 1649602621,
                },
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Comment',
          id: '30977197',
          text: 'I miss people like Erik in the Internet.<p>Erik was like a predator who contributed to the ecosystem. He ate frequent low quality shitposters from comp.lang.lisp and then swam in circles preventing them from emerging.<p>His rants were  educating or entertaining. Usually both. They came from deep technical knowledge. When was the last time you followed a Internet flamewar and learned something?',
          by: 'nabla9',
          time: 1649601081,
          parent: 30976616,
          comments: {
            edges: [
              {
                node: {
                  __typename: 'Comment',
                  id: '30977318',
                  text: 'It&#x27;s like the fondness of bullies in high school until you realize that other people are indeed people and not cardboard popups.',
                  by: 'throwmeariver1',
                  time: 1649601942,
                },
              },
            ],
          },
        },
      },
      {
        node: {
          __typename: 'Comment',
          id: '30976895',
          text: 'A reminder to get checked with the doctor regularly and to treat intestinal issues seriously.',
          by: 'mrtree',
          time: 1649598583,
          parent: 30976616,
          comments: {
            edges: [
              {
                node: {
                  __typename: 'Comment',
                  id: '30978859',
                  text: 'He had ulcerative colitis. People more or less use &quot;having an ulcer&quot; as slang for being in a bad mood. Maybe he was like the giant with a thorn in its foot - something bothered him, and that made him less than in a good mood.',
                  by: 'kzrdude',
                  time: 1649609184,
                },
              },
            ],
          },
        },
      },
    ],
  },
}
