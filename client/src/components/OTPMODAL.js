
import { useState } from "react";

export default function CenteredModal({ onClose, value, setValue, additional_message, setSubmit }) {
  const [code, setCode] = useState("");


  return (
    <div className="fixed z-50 inset-0 flex items-center justify-center bg-black/50 backdrop-blur-md">
      <div className="bg-white p-6 rounded-2xl shadow-lg w-96 text-center">

        <div id="additional-message" className="text-green-500">{additional_message}</div>
        <h2 className="text-xl font-bold">Enter Code</h2>
        <p className="text-gray-600 mt-2">Enter a 4-digit number</p>

        <input
          type="text"
          value={value}
          onChange={setValue}
          maxLength="4"
          className="mt-4 w-24 text-center border border-gray-300 rounded-lg p-2 text-xl outline-none focus:border-blue-500"
          placeholder="XXXX"
        />

        <div className="mt-4 flex justify-center space-x-4">
          <button
            onClick={onClose}
            className="px-4 py-2 bg-gray-400 text-white rounded-lg"
          >
            Cancel
          </button>
          <button onClick={setSubmit} className="px-4 py-2 bg-blue-500 text-white rounded-lg">
            Submit
          </button>
        </div>
      </div>
    </div>
  );
}
