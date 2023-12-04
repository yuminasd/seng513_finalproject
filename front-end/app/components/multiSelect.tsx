import React from 'react';

import Select from 'react-select';

const MultiSelect = ({ selectedGenres, onChange, darkMode }) => {
    const genreOptions = [
        'Canadian', 'Comedy Movies', 'TV Comedies', 'Mockumentaries', 'Late Night Comedy Movies', 'Stand-Up Comedy',
        'British', "Kids' TV", 'Kids & Family Movies', 'Kids Music', 'TV Cartoons', 'Kids Movies', 'Music',
        // ... (add the rest of your genres)
    ].map(genre => ({ value: genre, label: genre }));
    const selectStyles = {
        control: base => ({
            ...base,
            background: darkMode ? '#333' : '#fff', // Adjust background color
            color: darkMode ? '#fff' : '#333', // Adjust text color
        }),
        option: (provided, state) => ({
            ...provided,
            backgroundColor: state.isSelected ? (darkMode ? '#555' : '#ccc') : 'transparent', // Adjust selected option background color
            color: darkMode ? '#fff' : '#333', // Adjust selected option text color
        }),
        menu: base => ({
            ...base,
            backgroundColor: darkMode ? '#333' : '#fff', // Adjust dropdown menu background color
        }),
        menuList: base => ({
            ...base,
            color: darkMode ? '#fff' : '#333', // Adjust dropdown menu text color
        }),
    };
    return (<Select
        isMulti
        name="genres"
        options={genreOptions}
        value={selectedGenres}
        onChange={onChange}
        styles={selectStyles} // Apply custom styles
        className="basic-multi-select"
        classNamePrefix="select"
    />
    );
};

export default MultiSelect