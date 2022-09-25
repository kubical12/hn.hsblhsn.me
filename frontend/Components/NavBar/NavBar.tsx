import { Block } from 'baseui/block'
import {
  ALIGN,
  StyledNavigationItem,
  StyledNavigationList,
} from 'baseui/header-navigation'
import { Link } from 'react-router-dom'
import { useStyletron } from 'baseui'
import { Avatar } from 'baseui/avatar'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface NavBarProps {}

const NavBar: React.FC<NavBarProps> = () => {
  const [, theme] = useStyletron()
  return (
    <Block
      $style={{
        backgroundColor: theme.colors.backgroundPrimary,
      }}
    >
      <div className="flex items-left justify-left pl-2 pt-2 md:hidden">
        <Link to="/">
          <Logo />
        </Link>
      </div>
      <Block
        $style={{
          display: 'flex',
          borderBottomWidth: '1px',
          borderBottomStyle: 'solid',
          borderBottomColor: `${theme.colors.borderOpaque}`,
          fontSize: '1rem',
        }}
      >
        <StyledNavigationList $align={ALIGN.left}>
          <StyledNavigationItem className="mr-48 py-4 hidden md:block">
            <Link to="/">
              <Logo />
            </Link>
          </StyledNavigationItem>
          <NavBarItem to={'/'}>Top</NavBarItem>
          <NavBarItem to={'/newest'}>New</NavBarItem>
          <NavBarItem to={'/ask'}>Ask</NavBarItem>
          <NavBarItem to={'/show'}>Show</NavBarItem>
          <NavBarItem to={'/jobs'}>Jobs</NavBarItem>
          <NavBarItem to={'/search'}>Search</NavBarItem>
        </StyledNavigationList>
        <StyledNavigationList $align={ALIGN.center}></StyledNavigationList>
        <StyledNavigationList $align={ALIGN.right}>
          <StyledNavigationItem className="hidden md:block mr-12 cursor-not-allowed">
            <Avatar
              name="Coming Soon"
              size="scale900"
              src="https://avatars.dicebear.com/api/personas/hackernews.svg"
            />
          </StyledNavigationItem>
        </StyledNavigationList>
      </Block>
    </Block>
  )
}

interface NavBarItemProps {
  to: string
  children: React.ReactNode
}

const NavBarItem: React.FC<NavBarItemProps> = (props: NavBarItemProps) => {
  const [, theme] = useStyletron()
  //className="py-4 md:px-6 border-b-2 border-transparent md:hover:border-orange-500"
  return (
    <Link to={props.to}>
      <StyledNavigationItem
        $style={{
          color: theme.colors.contentPrimary,
          cursor: 'pointer',
          ':hover': {
            borderBottomWidth: '2px',
            borderBottomStyle: 'solid',
            borderBottomColor: theme.colors.accent ?? 'orange',
          },
          paddingTop: '1rem',
          paddingBottom: '1rem',
          borderBottomWidth: '2px',
          borderBottomStyle: 'solid',
          borderColor: 'transparent',
          paddingLeft: '0.7rem',
          paddingRight: '0.7rem',
          [theme.mediaQuery.medium]: {
            paddingLeft: '2rem',
            paddingRight: '2rem',
          },
        }}
      >
        {props.children}
      </StyledNavigationItem>
    </Link>
  )
}

const Logo: React.FC = () => {
  const [, theme] = useStyletron()
  return (
    <Block
      $style={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
      }}
    >
      <span
        className="flex-1 h-6 w-6 text-center mx-2"
        style={{
          backgroundColor: theme.colors.accent,
        }}
      >
        H
      </span>
      <span>HackerNews</span>
    </Block>
  )
}

export { NavBar }
export type { NavBarProps }
