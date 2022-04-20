import { CommentConnection } from './Comment'
import { User } from './User'

export interface Poll {
  by: User
  descendants: number
  id: string
  kids: number[]
  parts: number[]
  score: number
  text: string
  time: number
  title: string
  type: string
  comments: CommentConnection
}
