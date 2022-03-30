import {
  HeaderNavigation,
  ALIGN,
  StyledNavigationList,
  StyledNavigationItem,
} from 'baseui/header-navigation'
import { Block } from 'baseui/block'
import { HeadingSmall } from 'baseui/typography'
import { ArrowLeft } from 'baseui/icon'
import { FEED_KIND } from '../../hooks/navigation'
import { ArrowRight } from 'baseui/icon'
import { NavBtn } from './navBtn'
import { Link } from 'react-router-dom'
import { Loader } from './loader'

type NavbarPropsT = {
  onBack?: () => void
  onForward?: () => void
  feedKind?: FEED_KIND
  isLoading?: boolean
}

export function Navbar(props: NavbarPropsT) {
  const { onBack, onForward, isLoading } = props

  const BackBtn = (
    <NavBtn onClick={onBack}>
      <ArrowLeft />
    </NavBtn>
  )

  const ForwardBtn = (
    <NavBtn onClick={onForward}>
      <ArrowRight />
    </NavBtn>
  )

  return (
    <Block>
      <Loader isLoading={isLoading} />
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
          <StyledNavigationItem>{BackBtn}</StyledNavigationItem>
          <StyledNavigationItem>{ForwardBtn}</StyledNavigationItem>
          <StyledNavigationItem>
            <HeadingSmall>
              <Link to="/">HackerNews</Link>
            </HeadingSmall>
          </StyledNavigationItem>
        </StyledNavigationList>
      </HeaderNavigation>
      <Block
        $style={{
          marginBottom: '3.8rem',
        }}
      />
    </Block>
  )
}
