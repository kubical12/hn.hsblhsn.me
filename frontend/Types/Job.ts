import { OpenGraph } from './OpenGraph'
import { User } from './User'

export interface Job {
  id: string
  by: User
  score: number
  text: string
  time: number
  title: string
  type: string
  url: string
  openGraph?: OpenGraph
  html?: string
}
