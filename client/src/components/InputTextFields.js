import React from "react";

const InputTextField = ({ label, id, state = { state: "", setState: () => { } } }) => {
  return (
    <div className="mb-5">
      <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor={id}>
        {label}
      </label>
      <input
        id={id}
        value={state.state}
        onChange={(e) => state.setState(e.target.value)}
        type="text"
        className="w-full px-4 py-3 rounded-lg bg-gray-200 mt-2 border focus:border-blue-500 focus:bg-white focus:outline-none"
      />
    </div>
  );
};

export default InputTextField;
