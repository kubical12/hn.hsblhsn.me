import {
  HeaderNavigation,
  ALIGN,
  StyledNavigationList,
  StyledNavigationItem,
} from 'baseui/header-navigation'
import { Block } from 'baseui/block'
import LoadingBar from 'react-top-loading-bar'
import { HeadingSmall } from 'baseui/typography'
import ArrowLeft from 'baseui/icon/arrow-left'
import { FEED_KIND } from '../../hooks/navigation'
import { ArrowRight } from 'baseui/icon'
import { NavBtn } from './navBtn'
import { useEffect, useRef } from 'react'
import { useStyletron } from 'baseui'
import { Link } from 'react-router-dom'

type NavbarPropsT = {
  onBack?: () => void
  onForward?: () => void
  feedKind?: FEED_KIND
  isLoading?: boolean
}

export function Navbar(props: NavbarPropsT) {
  const ref = useRef(null)
  const [, theme] = useStyletron()
  const { onBack, onForward, isLoading } = props

  useEffect(() => {
    if (isLoading && isLoading === true) {
      ref.current.continuousStart()
    } else {
      ref.current.complete()
    }
  }, [isLoading])

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
      <LoadingBar color={theme.colors.accent} ref={ref} shadow={true} />
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
          marginBottom: '5rem',
        }}
      />
    </Block>
  )
}
