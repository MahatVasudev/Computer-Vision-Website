import { local_post } from "../api/newapi"

export const getPosts = (image_id) => {
  return `${local_post}/${image_id}`
}
