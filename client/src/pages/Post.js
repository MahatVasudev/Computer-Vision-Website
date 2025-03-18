import React, { useEffect, useState } from 'react';
import ImageWithModal from '../components/ImageWithModal';
import PostInfo from '../components/PostInfo';
import PostStats from '../components/PostStats';
import Comment from '../components/Comment';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPaperPlane } from '@fortawesome/free-solid-svg-icons';
import MasonryGrid from '../components/Masonry_Grid';
import { useParams, useSearchParams } from 'react-router-dom';
import { Pen } from 'lucide-react';
import { useGetPostDetailsMutation } from '../features/post/posts';
import { getPosts } from '../utils/common';

const Post = () => {
  const images = [
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRRCZVKWKAUmqHUszu8_M3CoepdRNIXk9SvZQ&s",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSNxnwEFPxa4ujQwSb3ebJQ0qScgFz7CEY14g&s",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTTrQSQC8BQOCvipNCu7-MdwTbVjqfASDE0pA&s",
    "https://wallpapers.com/images/hd/pink-aesthetic-anime-phone-zrwdo7l6d7dshf6k.jpg",
    "https://mfiles.alphacoders.com/101/thumb-1920-1013619.png",
    "https://lh3.googleusercontent.com/ZX6QtuP7xer0M5-ov_vw7K4qgjER0j_CnV5XoQ2KZU6DD2F4eRg2OyCeknEjvO-Uoyjk7_x5ljPJQe1e7F85aAdgMQ=s1280-w1280-h800",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS3ddrAtZQNd84Q2Ka_HBFozDlf81m2N7bEHg&s",
    "https://images7.alphacoders.com/135/thumb-1920-1352330.jpeg",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSCMQrAsyMardvp4iH8tqoypTJJHtZ9e5jfVg&s",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRfwEyltdnZT7_9BDyPtgUW8aXhdq5R3AkRqw&s",
  ];
  const [title, setTitle] = useState("")
  const [description, setDescription] = useState("")
  const [user_id, setUserId] = useState("")
  const [username, setUsername] = useState("")
  const [createdat, setCreatedat] = useState("")
  const [isPostLiked, setIsPostLiked] = useState(false);
  const [isPostBookmarked, setIsPostBookmarked] = useState(false);
  const [image, setImages] = useState("")
  const [getdetails, { isLoading }] = useGetPostDetailsMutation()

  const togglePostLike = () => {
    setIsPostLiked(!isPostLiked);
  };

  const { post_id } = useParams()
  const togglePostBookmark = () => {
    setIsPostBookmarked(!isPostBookmarked);
  };



  const GetPostDetails = async () => {
    try {
      const res = await getdetails(post_id)

      if (res.error) {
        console.log(res.error.data)
        return
      }

      setTitle(res.data.data.post_data.post_title)
      setDescription(res.data.data.post_data.post_desc)
      setUserId(res.data.data.post_data.user_id)
      setUsername(res.data.data.post_data.username)
      setCreatedat(res.data.data.post_data.createdat)
      setImages(res.data.data.post_data.images)
    } catch (e) {
      console.log(e)
    }
  }

  useEffect(() => {
    GetPostDetails()
  }, [])

  if (isLoading) {
    return (
      <>
        Loading

      </>
    )
  }
  return (
    <>
      <div className="flex w-full h-fit mt-[-1rem] overflow-hidden justify-center pt-4 pb-3 px-8 bg-transparent">

        <ImageWithModal
          imageUrl={getPosts(image)}
        />

        {/* Post details section */}
        <div className="w-1/3 ml-8 h-[44rem] bg-white shadow-lg rounded-lg p-6 flex flex-col overflow-hidden">

          {/* Post info and comments grow to fill available space */}
          <div className="flex-grow">
            <PostInfo
              creator={username}
              shortDescription={description}
              fullDescription={description}
              title={title}
            />

            {/* Comments section */}
            <Comment />
          </div>

          {/* Post stats and comment input stay at the bottom */}
          <div className="mt-[-10rem]">
            <PostStats
              isPostLiked={isPostLiked}
              isPostBookmarked={isPostBookmarked}
              toggleLike={togglePostLike}
              toggleBookmark={togglePostBookmark}
              imageUrl={getPosts(image)}
            />

            <p className="text-sm text-gray-500 mb-2">October 23rd 2004</p>

            {/* Comment input */}
            <div className="relative">
              <input
                type="text"
                placeholder="Comment Something Nice..."
                className="w-full border rounded-full py-2 px-4 text-gray-700 focus:outline-none"
              />
              <button className="absolute right-4 top-2 text-blue-500">
                <FontAwesomeIcon icon={faPaperPlane} />
              </button>
            </div>
          </div>
        </div>
      </div>

      <div className='text-3xl pt-3 mb-10 font-bold text-black'>Recommended Images</div>

      <MasonryGrid data={images} />
    </>
  );
};

export default Post;
