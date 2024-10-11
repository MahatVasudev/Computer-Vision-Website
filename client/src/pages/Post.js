import React, { useState } from 'react';
import { faHeart, faComment, faBookmark, faPaperPlane, faDownload } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const Post = () => {
  // State to track if the post has been liked
  const [isPostLiked, setIsPostLiked] = useState(false);

  // State to track if the post has been bookmarked
  const [isPostBookmarked, setIsPostBookmarked] = useState(false);

  // State to track if each comment has been liked
  const [likedComments, setLikedComments] = useState([false, false, false, false]);

  // State to track if each comment is expanded
  const [expandedComments, setExpandedComments] = useState([false, false, false, false]);

  // State to toggle the full description for the post
  const [isDescriptionExpanded, setIsDescriptionExpanded] = useState(false);

  // Function to toggle post like
  const togglePostLike = () => {
    setIsPostLiked(!isPostLiked);
  };

  // Function to toggle post bookmark
  const togglePostBookmark = () => {
    setIsPostBookmarked(!isPostBookmarked);
  };

  // Function to toggle like for a specific comment
  const toggleCommentLike = (index) => {
    const updatedLikes = [...likedComments];
    updatedLikes[index] = !updatedLikes[index]; // Toggle the liked state
    setLikedComments(updatedLikes);
  };

  // Function to toggle expand for a specific comment
  const toggleCommentExpand = (index) => {
    const updatedExpanded = [...expandedComments];
    updatedExpanded[index] = !updatedExpanded[index]; // Toggle the expanded state
    setExpandedComments(updatedExpanded);
  };

  // Toggle the description view
  const toggleDescription = () => {
    setIsDescriptionExpanded(!isDescriptionExpanded);
  };

  // Short description and full description
  const shortDescription = "Lorem ipsum dolor sit amet, consectetur adipiscing elit.";
  const fullDescription = `${shortDescription} Cras facilisis neque orci, et cursus lectus bibendum sed. Nam eget vehicula magna. Suspendisse potenti. Fusce volutpat sapien non est sodales tempus. Etiam vestibulum sapien eget eros blandit feugiat. Integer vel fringilla quam.`;

  // Limit for when the "more" button should appear for the post description
  const shouldShowMorePostDesc = fullDescription.length > shortDescription.length;

  return (
    <div className="flex w-full h-screen overflow-hidden p-8 bg-gray-100">
      {/* Left: Image Section */}
      <div className="w-2/3 bg-white shadow-lg rounded-lg flex justify-center items-center">
        <img
          src="https://via.placeholder.com/300x500/8A2BE2/FFFFFF?text=Post+Image" // Replace this with the actual image URL
          alt="Post"
          className="object-cover h-full rounded-lg"
        />
      </div>

      {/* Right: Post Details Section */}
      <div className="w-1/3 ml-8 bg-white shadow-lg rounded-lg p-6 flex flex-col overflow-hidden"> {/* Added overflow-hidden */}
        {/* Post Title */}
        <h1 className="text-2xl font-bold mb-4">Title</h1>

        {/* Creator Info */}
        <div className="flex items-center mb-4">
          <img
            src="/user_default.jpg" // Replace with creator's profile image
            alt="Creator Avatar"
            className="h-10 w-10 rounded-full border"
          />
          <div className="ml-2">
            <p className="font-semibold">@creator</p>
            <p className="text-sm text-gray-500">
              {isDescriptionExpanded ? fullDescription : `${shortDescription} `}
              {shouldShowMorePostDesc && (
                <span
                  className="text-blue-500 cursor-pointer"
                  onClick={toggleDescription}
                >
                  {isDescriptionExpanded ? "Show less" : "...more"}
                </span>
              )}
            </p>
          </div>
        </div>

        {/* Comment Section */}
        <div className="flex-grow overflow-y-auto mb-4">
          {['commenter1', 'commenter2', 'commenter3', 'commenter4'].map((commenter, index) => {
            const shortComment = "This is a comment.";
            const fullComment = `${shortComment} More details about the comment. More details.`;
            const isCommentExpanded = expandedComments[index];
            const shouldShowMoreComment = fullComment.length > shortComment.length;

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
                    <p className="text-sm text-gray-500">
                      {isCommentExpanded ? fullComment : `${shortComment} `}
                      {shouldShowMoreComment && (
                        <span
                          className="text-blue-500 cursor-pointer"
                          onClick={() => toggleCommentExpand(index)} // Only toggle expand for the comment
                        >
                          {isCommentExpanded ? "Show less" : "...more"}
                        </span>
                      )}
                    </p>
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

        {/* Post Stats */}
        <div className="border-t py-4">
          <div className="flex items-center justify-between">
            {/* Left section: Like, Comment */}
            <div className="flex space-x-4">
              <FontAwesomeIcon
                icon={faHeart}
                className={`cursor-pointer ${isPostLiked ? 'text-red-500' : 'text-gray-600'}`}
                onClick={togglePostLike}
              />
              <FontAwesomeIcon icon={faComment} className="text-gray-600" />
            </div>
            {/* Right section: Download, Bookmark */}
            <div className="ml-auto flex space-x-4">
              <FontAwesomeIcon icon={faDownload} className="text-gray-600" />
              <FontAwesomeIcon
                icon={faBookmark}
                className={`cursor-pointer ${isPostBookmarked ? 'text-yellow-500' : 'text-gray-600'}`}
                onClick={togglePostBookmark}
              />
            </div>
          </div>

          {/* Post Stats - now on a new line */}
          <p className="text-sm text-gray-500 mt-2">10 Likes · 3 Comments · 2 Saves</p>
        </div>

        {/* Date */}
        <p className="text-sm text-gray-500 mb-2">October 23rd 2004</p>

        {/* Add Comment Input */}
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
  );
};

export default Post;
