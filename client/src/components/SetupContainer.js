import React from "react";
import { ArrowLeft } from "lucide-react";

const CenteredContainer = ({ children, prev_fn, next_fn }) => {
  return (
    <div className="flex flex-col items-center justify-center h-screen/50 bg-white p-4">
      {/* Title */}
      <h1 className="text-3xl font-bold mb-4 text-center">
        Tell us more about yourself
      </h1>

      {/* Card Container */}
      <div className="relative bg-gray-100 border-2 w-[80rem] min-h-[35rem] max-w-full rounded-2xl p-6 flex flex-col">
        {/* Back Button */}
        <button className="absolute top-4 left-4 bg-gray-400 rounded-full p-2" onClick={prev_fn}>
          <ArrowLeft size={20} color="black" />
        </button>

        {/* Dynamic Content */}
        <div className="flex flex-col p-5 mt-5 mb-auto justify-center overflow-auto">{children}</div>

        {/* Next Button */}
        <button
          onClick={next_fn}
          className="absolute bottom-4 right-4 bg-blue-500 text-white font-bold rounded-full px-4 py-2">
          Next
        </button>
      </div>
    </div>
  );
};

export default CenteredContainer;
