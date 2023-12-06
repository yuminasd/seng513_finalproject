import React from 'react';

interface ButtonProps {
    text: string;
    color: 'primary' | 'secondary'; // Define accepted color values
    onClick?: () => void; // Optional onClick event
}

const Button: React.FC<ButtonProps> = ({ text, color, onClick }) => {
    const buttonClasses = `w-full py-2 px-4 rounded-md ${color === 'primary' ? 'bg-purple-500 text-white' : 'border border-neutral-500 text-white'
        }`;

    return (
        <button className={buttonClasses} onClick={onClick}>
            {text}
        </button>
    );
};

export default Button;