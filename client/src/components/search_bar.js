import { React } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";

const SearchBar = ({ top_style, bottom_style, button_style, theme }) => {
	return (
		<div className={`flex items-center bg-[#D9D9D9] rounded-full px-4 py-2 w-[40rem] ${top_style}`}>
			<input
				type="text"
				placeholder="Search"
				className={`bg-transparent text-black placeholder-black outline-none w-full ${bottom_style}`}
			/>
			<FontAwesomeIcon icon={faSearch} className={`w-5 h-5 ${button_style}`} />
		</div>
	)
}

export default SearchBar
