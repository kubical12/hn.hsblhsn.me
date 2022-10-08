import { User } from '../../Types/User'
import { StatefulPopover } from 'baseui/popover'
import { Card, StyledBody } from 'baseui/card'
import {
  HeadingXSmall,
  ParagraphSmall,
  ParagraphXSmall,
} from 'baseui/typography'
import { useStyletron } from 'baseui'
import { fromNow } from '../commonutils'
import { Fragment } from 'react'

interface UserPopoverProps {
  user: User
  children: React.ReactNode
}

const UserPopover: React.FC<UserPopoverProps> = ({
  user,
  children,
}: UserPopoverProps) => {
  const [, theme] = useStyletron()
  if (!user.id) {
    return <Fragment>{children}</Fragment>
  }
  return (
    <StatefulPopover
      triggerType="hover"
      content={() => (
        <Card
          overrides={{
            Root: {
              style: {
                maxWidth: '280px',
                backgroundColor: theme.colors.backgroundSecondary,
              },
            },
          }}
        >
          <Fragment>
            <HeadingXSmall color={theme.colors.accent}>{user.id}</HeadingXSmall>
            <StyledBody>
              <ParagraphXSmall>
                <span style={{ color: theme.colors.contentTertiary }}>
                  Joined&nbsp;{fromNow((user.created || 0) * 1000)} ago
                  &nbsp;&middot;&nbsp;{user.karma} Karma
                </span>
              </ParagraphXSmall>
              <ParagraphXSmall>
                <span style={{ color: theme.colors.contentTertiary }}>
                  {user.submitted?.totalCount || 0} submissions
                </span>
              </ParagraphXSmall>
              <br />
              <ParagraphSmall>
                <span dangerouslySetInnerHTML={{ __html: user.about || '' }} />
              </ParagraphSmall>
            </StyledBody>
          </Fragment>
        </Card>
      )}
    >
      {children}
    </StatefulPopover>
  )
}

export { UserPopover }
