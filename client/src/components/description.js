import React from 'react'

import { useState } from "react";

const Description = ({ full_comment }) => {
  const [expandedComments, setExpandedComments] = useState([false]);
  const shortComment = "This is a comment.";
  const fullComment = `${shortComment} More details about the comment. More details.`;
  const shouldShowMoreComment = fullComment.length > shortComment.length;
  const isCommentExpanded = expandedComments[1];
  const toggleCommentExpand = (key) => {
    const updatedExpanded = [...expandedComments];
    updatedExpanded[1] = !updatedExpanded[1]; // Toggle the expanded state
    setExpandedComments(updatedExpanded);
  };



  return (
    <>
      <p className="text-sm text-gray-500">
        {isCommentExpanded ? fullComment : `${shortComment} `}
        {shouldShowMoreComment && (
          <span
            className="text-blue-500 cursor-pointer"
            onClick={() => toggleCommentExpand(1)} // Only toggle expand for the comment
          >
            {isCommentExpanded ? "Show less" : "...more"}
          </span>
        )}
      </p>

    </>
  )
}

export default Description
