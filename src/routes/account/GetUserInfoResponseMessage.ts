import type { User } from "./User";

export interface GetUserInfoResponseMessage {
  id: number;
  title: string;
  date: number;
  msg: string;
  users: [User];
}
