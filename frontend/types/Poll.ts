import { CommentConnection } from "./Comment"

export interface Poll {
  by: string
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
