
import React from 'react';

const EditOptions = () => {
  return (
    <div className="space-y-4 mb-4">
      <div>
        <input type="checkbox" id="autoColorization" />
        <label htmlFor="autoColorization" className="ml-2">Auto Colorization</label>
      </div>
      <div>
        <input type="checkbox" id="autoBlur" />
        <label htmlFor="autoBlur" className="ml-2">Auto Background Blur</label>
      </div>
    </div>
  );
};

export default EditOptions;
