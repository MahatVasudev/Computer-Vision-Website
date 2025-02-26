import React, { useState } from "react";
import Description from "./description";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHeart } from "@fortawesome/free-solid-svg-icons";

const Comment = () => {
  const [likedComments, setLikedComments] = useState([false, false, false, false]);

  const toggleCommentLike = (index) => {
    const updatedLikes = [...likedComments];
    updatedLikes[index] = !updatedLikes[index]; // Toggle the liked state
    setLikedComments(updatedLikes);
  };

  return (
    <div className="flex-grow overflow-y-scroll max-h-96 mb-4">
      {['commenter', 'commenter1', 'commenter', 'commenter2', 'commenter3', 'commenter4', 'commenter', "commenter"].map((commenter, index) => {
        return (
          <div key={index} className="flex items-center justify-between mb-4">
            <div className="flex items-center">
              <img
                src="/user_default.jpg" // Replace with commenter's profile image
                alt="Commenter Avatar"
                className="h-8 w-8 rounded-full border"
              />
              <div className="ml-2">
                <p className="font-semibold">@{commenter}</p>
                <Description full_comment={"Something"} />
              </div>
            </div>
            {/* Comment like icon */}
            <div className="flex items-center">
              <FontAwesomeIcon
                icon={faHeart}
                className={`cursor-pointer ${likedComments[index] ? 'text-red-500' : 'text-gray-500'}`}
                onClick={() => toggleCommentLike(index)} // Only toggle like for the comment
              />
              <span className="ml-1 text-sm">10</span>
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default Comment;
