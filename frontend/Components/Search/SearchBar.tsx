import { useStyletron } from 'baseui'
import { FormEvent, useCallback } from 'react'
import { Block } from 'baseui/block'
import { HeadingXXLarge } from 'baseui/typography'
import { FormControl } from 'baseui/form-control'
import { Input } from 'baseui/input'
import { Search } from 'baseui/icon'

interface SearchBarProps {
  value: string
  onChange: (val: string) => void
}

const SearchBar = ({ value, onChange }: SearchBarProps) => {
  const [css, theme] = useStyletron()
  const focusOut = useCallback((e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const el = document.getElementById('search-input')
    if (el) {
      el.blur()
    }
    return false
  }, [])
  return (
    <Block
      className={css({
        textAlign: 'center',
        paddingTop: theme.sizing.scale300,
      })}
    >
      <HeadingXXLarge
        className={css({
          textAlign: 'center',
          paddingBottom: theme.sizing.scale600,
        })}
      >
        Search HackerNews!
      </HeadingXXLarge>
      <form onSubmit={focusOut} onReset={focusOut}>
        <FormControl
          //label="Search HackerNews"
          caption="All search results are sorted by relevance."
        >
          <Input
            id="search-input"
            name="q"
            type="search"
            clearable={true}
            placeholder="Type to search..."
            value={value}
            onChange={(e) => onChange(e.currentTarget.value)}
            autoFocus={true}
            startEnhancer={<Search size={24} />}
          />
        </FormControl>
      </form>
    </Block>
  )
}

export { SearchBar }
