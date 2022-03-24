import {
  HeaderNavigation,
  ALIGN,
  StyledNavigationList,
  StyledNavigationItem,
} from 'baseui/header-navigation'
import { Block } from 'baseui/block'
import { Link } from 'react-router-dom'
import { HeadingLarge } from 'baseui/typography'

export function Navbar() {
  return (
    <Block>
      <HeaderNavigation
        overrides={{
          Root: {
            style: ({ $theme }) => ({
              backgroundColor: $theme.colors.backgroundPrimary,
              position: 'fixed',
              top: 0,
              left: 0,
              right: 0,
              zIndex: 10,
            }),
          },
        }}
      >
        <StyledNavigationList $align={ALIGN.left}>
          <StyledNavigationItem>
            <HeadingLarge>
              <Link to="/" className="sm:hidden">
                HN
              </Link>
              <Link to="/" className="hidden sm:block">
                HackerNews
              </Link>
            </HeadingLarge>
          </StyledNavigationItem>
        </StyledNavigationList>
      </HeaderNavigation>
      <Block className="mb-16"></Block>
    </Block>
  )
}
