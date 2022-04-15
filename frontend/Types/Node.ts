import { Comment } from './Comment'
import { Job } from './Job'
import { PageInfo } from './PageInfo'
import { Poll } from './Poll'
import { PollOption } from './PollOption'
import { Story } from './Story'

export type Item = Story | Job | Comment | Poll | PollOption
export type NodeT<T> = T
export interface EdgeT<T> {
  node: NodeT<T>
}
export interface ConnectionT<T> {
  edges: EdgeT<T>[]
  pageInfo: PageInfo
}
