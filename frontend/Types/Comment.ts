import { ConnectionT } from './Node'
import { User } from './User'

export interface Comment {
  by: User
  id: string
  kids: number[]
  parent: number
  text: string
  time: number
  type: string
  dead: boolean
  deleted: boolean
  comments: CommentConnection
}

export type CommentConnection = ConnectionT<Comment>
