import { useStyletron } from 'baseui'
import { Block } from 'baseui/block'
import { StyledLink } from 'baseui/link'
import {
  HeadingLarge,
  HeadingSmall,
  Label4,
  Paragraph4,
} from 'baseui/typography'
import moment from 'moment'
import { Link } from 'react-router-dom'
import { Job, NodeT, Story } from '../../types'
import { CommentThread } from '../CommentThread'
import { getHost, getLink } from '../commonutils'
import './Item.css'

interface ItemProps {
  item: NodeT<Story | Job>
}

const Item: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()

  if (!item || (item.type !== 'story' && item.type !== 'job')) {
    // eslint-disable-next-line unicorn/no-null
    return null
  }
  return (
    <Block paddingTop={theme.sizing.scale600}>
      <Header item={item} />
      <Content item={item} />
      <Comments item={item} />
      <ContentLinks item={item} />
    </Block>
  )
}

const Header: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [css, theme] = useStyletron()
  const color = (c: string) =>
    css({
      color: c,
      cursor: 'pointer',
    })
  return (
    <Block>
      <Label4>
        <span className={color(theme.colors.accent)}>@{item.by}</span>&nbsp;
        <span className={color(theme.colors.colorSecondary)}>
          {item.time ? moment(item.time * 1000).fromNow() : 'unknown'}
        </span>
      </Label4>
      <HeadingLarge>
        <Link to={`/items?id=${item.id}`}>
          {item.openGraph?.title || item.title}
        </Link>
      </HeadingLarge>
      <Label4
        color={theme.colors.contentTertiary}
        paddingBottom={theme.sizing.scale800}
      >
        Read on&nbsp;
        <StyledLink
          href={getLink(item.id, item.url)}
          target="_blank"
          rel="noreferrer"
        >
          {getHost(item.id, item.url)}
        </StyledLink>
      </Label4>
    </Block>
  )
}

const Content: React.FC<ItemProps> = ({ item }: ItemProps) => {
  let val = ''
  if ('text' in item && item.text !== '') {
    val = item.text
  } else if ('html' in item && item.html !== '') {
    val = item.html || ''
  } else if (item.openGraph?.description) {
    val = item.openGraph.description
  }
  return (
    <section
      id="reader-view-content"
      dangerouslySetInnerHTML={{ __html: val }}
    />
  )
}

const Comments: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()
  return (
    <Block>
      {'descendants' in item && (
        <HeadingSmall paddingTop={theme.sizing.scale1200}>
          {item.descendants} comments
          <Label4 color={theme.colors.contentTertiary}>
            Posted on&nbsp;
            <StyledLink
              href={getLink(item.id, undefined)}
              target="_blank"
              rel="noreferrer"
            >
              {getHost(item.id, undefined)}
            </StyledLink>
          </Label4>
        </HeadingSmall>
      )}
      {'comments' in item && (
        <Block paddingBottom={theme.sizing.scale2400}>
          <CommentThread comments={item.comments} />
        </Block>
      )}
    </Block>
  )
}

const ContentLinks: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()
  return (
    <Block>
      <Paragraph4
        as="div"
        paddingTop={theme.sizing.scale2400}
        paddingBottom={theme.sizing.scale2400}
      >
        <Label4>
          Contents loaded from&nbsp;&nbsp;
          <StyledLink href={getLink(item.id, item.url)}>
            {getLink(item.id, item.url)}
          </StyledLink>
        </Label4>
        <Label4>
          Comments loaded from&nbsp;&nbsp;
          <StyledLink href={getLink(item.id, undefined)}>
            {getLink(item.id, undefined)}
          </StyledLink>
        </Label4>
      </Paragraph4>
    </Block>
  )
}

export { Item }
export type { ItemProps }
