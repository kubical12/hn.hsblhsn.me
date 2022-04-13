import { ConnectionT } from './Node'

export interface Comment {
  by: string
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
