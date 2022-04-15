import { OpenGraph } from "./OpenGraph"

export interface Job {
  id: string
  by: string
  score: number
  text: string
  time: number
  title: string
  type: string
  url: string
  openGraph?: OpenGraph
  html?: string
}
