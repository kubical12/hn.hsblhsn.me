import { Block } from 'baseui/block'
import {
  ALIGN,
  StyledNavigationItem,
  StyledNavigationList,
} from 'baseui/header-navigation'
import { Link, useLocation } from 'react-router-dom'
import { useStyletron } from 'baseui'
import { Avatar } from 'baseui/avatar'
import { Fragment, useEffect, useState } from 'react'
import { BEHAVIOR, Cell, Grid } from 'baseui/layout-grid'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface NavBarProps {}

const NavBar: React.FC<NavBarProps> = () => {
  const [, theme] = useStyletron()
  return (
    <Fragment>
      <Grid
        behavior={BEHAVIOR.fixed}
        gridColumns={[4, 8, 12]}
        gridMargins={[0, 0, 0]}
        gridGaps={0}
        gridMaxWidth={1200}
        overrides={{
          Grid: {
            style: {
              position: 'fixed',
              top: 0,
              backgroundColor: theme.colors.backgroundSecondary,
            },
          },
        }}
      >
        <Cell span={[4, 8, 12]}>
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
              <StyledNavigationItem className="hidden md:block py-4 mr-12 xl:mr-48">
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
        </Cell>
      </Grid>
      <Block
        $style={{
          height: '5rem',
          [theme.mediaQuery.medium]: {
            height: '4rem',
          },
        }}
      ></Block>
    </Fragment>
  )
}

interface NavBarItemProps {
  to: string
  children: React.ReactNode
}

const NavBarItem: React.FC<NavBarItemProps> = (props: NavBarItemProps) => {
  const [, theme] = useStyletron()
  const location = useLocation()
  const [isActive, setIsActive] = useState(location.pathname === props.to)
  useEffect(() => {
    setIsActive(location.pathname === props.to)
  }, [location.pathname])

  return (
    <Link to={props.to}>
      <StyledNavigationItem
        $style={{
          color: theme.colors.contentPrimary,
          cursor: 'pointer',
          ':hover': {
            borderBottomWidth: '2px',
            borderBottomStyle: 'solid',
            borderBottomColor: 'orange',
          },
          paddingTop: '1rem',
          paddingBottom: '1rem',
          borderBottomWidth: '2px',
          borderBottomStyle: 'solid',
          borderColor: isActive ? theme.colors.accent : 'transparent',
          paddingLeft: '0.7rem',
          paddingRight: '0.7rem',
          [theme.mediaQuery.medium]: {
            paddingLeft: '1rem',
            paddingRight: '1rem',
          },
          [theme.mediaQuery.large]: {
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
