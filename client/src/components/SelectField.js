
import { ChevronDown } from "lucide-react";
const SelectField = ({ options = [{ value: "", label: "" }], state = { state: "", setState: () => { } } }) => {
  console.log(options)
  return (
    <div className="relative w-64">
      <select
        className="appearance-none w-full px-4 py-3 text-lg bg-gray-900 text-white border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        value={state.state}
        onChange={(e) => state.setState(e.target.value)}
      >
        {
          options.map((e) =>

            (<option key={e.value} value={e.value} disabled={e.value === ""}>{e.label}</option>)
          )
        }
      </select>

      {/* Custom dropdown arrow */}
      <div className="absolute top-1/2 right-4 transform -translate-y-1/2 pointer-events-none">
        <ChevronDown size={20} className="text-white" />
      </div>
    </div>
  )
}

export default SelectField
