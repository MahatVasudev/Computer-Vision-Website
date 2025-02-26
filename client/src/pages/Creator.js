import React, { useEffect, useState } from 'react';
import MasonryGrid from '../components/Masonry_Grid'; // A separate component for the masonry grid
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useUserdetailsMutation } from '../features/user/user';
import { useNavigate, useParams } from 'react-router-dom';

const CreatorPage = () => {
  const [selectedTab, setSelectedTab] = useState('Posts'); // State for managing selected tab
  const [details, { isLoading }] = useUserdetailsMutation()
  const [err, setErr] = useState("")
  const [namesField, setNames] = useState({})
  const [secondaryDetails, setSecondaryDetails] = useState({})
  const [assests, SetAssets] = useState({})
  const [permisson, SetPermission] = useState([])
  const { creator } = useParams()
  const navigate = useNavigate()
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
      setNames({ first_name: user_data.first_name, last_name: user_data.last_name })
      setSecondaryDetails({ username: user_data.username })
      SetPermission(permissions)

      console.log(namesField)
    } catch (err) {
      setErr(err)
    }
  }

  useEffect(() => {
    getDetails()
  }, [])
  const imageUrls = [
    "https://via.placeholder.com/300x400/FF7F50/FFFFFF?text=Image+1",
    "https://via.placeholder.com/300x300/6495ED/FFFFFF?text=Image+2",
    "https://via.placeholder.com/300x500/FF69B4/FFFFFF?text=Image+3",
    "https://via.placeholder.com/300x350/8A2BE2/FFFFFF?text=Image+4",
    "https://via.placeholder.com/300x450/00FA9A/FFFFFF?text=Image+5",
    "https://via.placeholder.com/300x300/FFD700/FFFFFF?text=Image+6",
    "https://via.placeholder.com/300x400/DC143C/FFFFFF?text=Image+7",
    "https://via.placeholder.com/300x350/ADFF2F/FFFFFF?text=Image+8",
    "https://via.placeholder.com/300x500/20B2AA/FFFFFF?text=Image+9",
    "https://via.placeholder.com/300x400/FF7F50/FFFFFF?text=Image+10",
  ];

  if (isLoading) {
    return <>
      <div>Loading</div>
    </>
  }
  return (
    <>
      <div className="w-full flex flex-col items-center rounded-b-xl shadow-lg">
        {/* Cover Photo */}
        <img
          src="cover_image.jpg" // Replace with your actual cover photo URL
          alt="Cover"
          className="w-full h-48 object-cover" // Adjusted to make it a cover photo
        />

        {/* Creator Info */}
        <div className="mt-4 w-full px-8 flex items-center justify-between">
          {/* Left side: Avatar + Creator details */}
          <div className="flex items-center">
            {/* Avatar stays at the left */}
            <img
              src="/user_avatar.jpg"
              alt="Creator Avatar"
              className="h-[10rem] w-[10rem] rounded-full border-4 border-white relative -translate-y-[55%]" />


            <div className="ml-4 -translate-y-8  p-4">
              <h1 className="text-3xl font-bold">{namesField.first_name} {namesField.last_name}</h1>
              <p className="text-gray-600">@{secondaryDetails?.username} · AB Followers · AB Photos</p>
              <p className="mt-2 text-gray-500">Description... </p>
            </div>
          </div>

          {/* Follow button */}
          {
            permisson.includes("permission:edit") ?
              (<>

                <button className="px-6 py-2 bg-blue-500 text-white rounded-full flex items-center hover:bg-blue-600">
                  Edit
                </button>
              </>) :
              (<>

                <button className="px-6 py-2 bg-blue-500 text-white rounded-full flex items-center hover:bg-blue-600">
                  <FontAwesomeIcon icon={faPlus} className="mr-2" />
                  Follow
                </button>
              </>)
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
        <MasonryGrid data={imageUrls} />
      </div>
    </>
  );
};

export default CreatorPage;
