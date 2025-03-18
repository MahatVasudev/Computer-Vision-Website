import React, { useEffect, useState } from 'react';
import MasonryGrid from '../components/Masonry_Grid'; // A separate component for the masonry grid
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useGetAllUserPostsMutation, useUserdetailsMutation } from '../features/user/user';
import { useNavigate, useParams } from 'react-router-dom';
import { useFollowUserMutation } from '../features/follow/follow';

const CreatorPage = () => {
  const [selectedTab, setSelectedTab] = useState('Posts'); // State for managing selected tab
  const [details, { isLoading }] = useUserdetailsMutation()
  const [err, setErr] = useState("")
  const [namesField, setNames] = useState({})
  const [secondaryDetails, setSecondaryDetails] = useState({})
  const [permisson, SetPermission] = useState([])
  const { creator } = useParams()
  const [id, setId] = useState("")
  const navigate = useNavigate()
  const [followstatus, setFollowStatus] = useState({})
  const [follow] = useFollowUserMutation()
  const [posts, setPosts] = useState([])
  const [getposts, { }] = useGetAllUserPostsMutation()


  const getDetails = async () => {
    try {
      const data = await details({ username: creator })

      if (data.error) {
        if (data.error.status === 404) {
          navigate("/")
          return
        }

        setErr(data.error)
        return
      }
      const user_data = data.data.data.user

      const permissions = data.data.data.permissions
      const follow = data.data.data.follows
      setId(user_data.id)
      setNames({ first_name: user_data.first_name, last_name: user_data.last_name })
      setSecondaryDetails({ username: user_data.username, followings: follow.following, followers: follow.followers, post_count: data.data.data.post_count })
      setFollowStatus({ followed: data.data.data.isfollowing, followed_by: data.data.data.isfollowedby })
      SetPermission(permissions)

      console.log(namesField)
    } catch (err) {
      setErr(err)
    }
  }

  const FollowFunction = async () => {
    try {

      const response = await follow({
        "following_id": id,
        type: true
      })

      console.log(response)
      if (response.error) {
        if (response.error.code === 409) {
          setErr("Cannot Follow Yourself Self or Someone Twice")
          return
        } else if (response.error.code === 404) {
          setErr("User Not Found")
          return
        } else {
          setErr("Server Error: ", response.error.data)
          return
        }
      }



      setFollowStatus({ ...followstatus, followed: true })
      setSecondaryDetails(preDetails => ({ ...preDetails, followers: preDetails.followers + 1 }))
    } catch (errors) {
      console.log(errors)
    }
  }

  const GetAllPostsFromUser = async () => {
    try {
      const res = await getposts(creator)

      if (res.error) {
        alert(res.error.message)
        return
      }

      setPosts(res.data.data)

    } catch (e) {
      console.log(e)
    }
  }


  useEffect(() => {
    getDetails()
    GetAllPostsFromUser()
  }, [])

  if (isLoading) {
    return <>
      <div>Loading</div>
    </>
  }
  return (
    <div className='dark:dark:bg-[#272727]'>
      <div className="w-full flex flex-col items-center rounded-b-xl shadow-lg">
        {/* Cover Photo */}
        <img
          src="https://wallpapers.com/images/hd/pink-aesthetic-anime-phone-zrwdo7l6d7dshf6k.jpg" // Replace with your actual cover photo URL
          alt="Cover"
          className="w-full h-48 object-cover" // Adjusted to make it a cover photo
        />

        {/* Creator Info */}
        <div className="mt-4 w-full dark:bg-[#242424] px-8 flex items-center justify-between">
          {/* Left side: Avatar + Creator details */}
          <div className="flex items-center">
            {/* Avatar stays at the left */}
            <img
              src="/user_avatar.jpg"
              alt="Creator Avatar"
              className="h-[10rem] w-[10rem] rounded-full border-4 border-white relative -translate-y-[55%]" />


            <div className="ml-4 -translate-y-8 dark:text-white  p-4">
              <h1 className="text-3xl font-bold">{namesField.first_name} {namesField.last_name}</h1>
              <p className="text-gray-600 dark:text-white">@{secondaryDetails?.username} · {secondaryDetails?.followers} Followers · {secondaryDetails?.followings} Following · {secondaryDetails?.post_count} Posts</p>
              <p className="mt-2 text-gray-500 dark:text-white">Description... </p>
            </div>
          </div>

          {/* Follow button */}
          {
            permisson.includes("profile:edit") ? (
              <button

                className="px-6 py-2 bg-blue-500 text-white rounded-full flex items-center hover:bg-blue-600">
                Edit
              </button>
            ) : !followstatus.followed ? (
              <button
                onClick={FollowFunction}
                className="px-6 py-2 bg-blue-500 text-white rounded-full flex items-center hover:bg-blue-600"
              >
                <FontAwesomeIcon icon={faPlus} className="mr-2" />
                {followstatus.followed_by ? "Follow Back" : "Follow"}
              </button>
            ) : (
              <button
                className="px-6 py-2 bg-blue-500 text-white rounded-full flex items-center hover:bg-blue-600"
              >
                <FontAwesomeIcon icon={faPlus} className="mr-2" />
                Unfollow
              </button>
            )
          }
        </div>
      </div>

      {/* Tab Navigation */}
      <div className="w-full border-b mt-4">
        <nav className="flex text-2xl font-bold ml-6 space-x-8">
          {['Posts', 'Saved', 'Liked'].map(tab => (
            <button
              key={tab}
              onClick={() => setSelectedTab(tab)}
              className={`py-2 ${selectedTab === tab ? 'text-blue-600 border-b-2 border-blue-600' : 'text-gray-500'}`}
            >
              {tab}
            </button>
          ))}
        </nav>
      </div>

      {/* Masonry Grid Content */}
      <div className="lg:mx-[8rem] mx-[2rem]  mt-6">
        <MasonryGrid data={posts} />
      </div>
    </div>
  );
};

export default CreatorPage;
