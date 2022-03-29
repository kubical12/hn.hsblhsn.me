import {
  useNavigate,
  useLocation,
  useResolvedPath,
  matchPath,
} from 'react-router-dom'

export enum FEED_KIND {
  new = 'new',
  top = 'top',
}

export enum PAGE {
  home = '/',
  feed = '/feed/:kind',
  read = '/read/:id',
  notFound = '/not-found',
}

export type AppNavigatorT = {
  currentPage(): PAGE
  home(): void
  feed(kind: FEED_KIND, page: number): void
  read(id: number): void
}

const useAppNavigator = (): AppNavigatorT => {
  const navigate = useNavigate()
  const location = useLocation()
  const resolved = useResolvedPath(location.pathname)

  const feed = (kind: FEED_KIND, page: number) => {
    navigate(`/feed/${kind}?page=${page}`)
    return
  }
  const read = (id: number) => {
    navigate(`/read/${id}`)
    return
  }
  const home = () => {
    navigate('/')
    return
  }

  const currentPage = (): PAGE => {
    if (matchPath('/', resolved.pathname)) {
      return PAGE.home
    } else if (matchPath('/feed/:kind', resolved.pathname)) {
      return PAGE.feed
    } else if (matchPath('/read/:id', resolved.pathname)) {
      return PAGE.read
    } else {
      return PAGE.notFound
    }
  }
  return {
    home,
    currentPage,
    feed,
    read,
  }
}

export default useAppNavigator
