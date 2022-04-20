import { User } from './User'

export interface PollOption {
  by: User
  id: string
  parent: number
  score: number
  text: string
  time: number
  type: string
}
