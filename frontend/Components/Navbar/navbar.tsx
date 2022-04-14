import { Block } from 'baseui/block'
import {
  HeaderNavigation,
  StyledNavigationItem,
  StyledNavigationList,
  ALIGN,
} from 'baseui/header-navigation'
import { Link } from 'react-router-dom'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface NavBarProps {}

const NavBar: React.FC<NavBarProps> = () => {
  return (
    <Block>
      <HeaderNavigation
        overrides={{
          Root: {
            style: {
              backgroundColor: '#ff6600',
            },
          },
        }}
      >
        <StyledNavigationList $align={ALIGN.left}>
          <StyledNavigationItem className="hidden md:block">
            <Link to="/"> HackerNews</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/">Top</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/newest">New</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/ask">Ask</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/show">Show</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/jobs">Jobs</Link>
          </StyledNavigationItem>
        </StyledNavigationList>
      </HeaderNavigation>
    </Block>
  )
}

export { NavBar }
export type { NavBarProps }
