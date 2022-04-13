import { CommentConnection } from './Comment'
import { OpenGraph } from './Opengraph'

export interface Story {
  by: string
  descendants: number
  id: string
  kids: number[]
  score: number
  time: number
  title: string
  type: string
  url: string
  comments: CommentConnection
  openGraph?: OpenGraph
  html?: string
}
