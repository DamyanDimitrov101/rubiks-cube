import React from 'react';

const FaceButton = ({ face, label, direction, handleRotate, loading }) => {
    const isClockwise = direction === 'clockwise';
    const directionSymbol = isClockwise ? '↻' : '↺';

    return (
        <button
            onClick={() => handleRotate(face, isClockwise)}
            disabled={loading}
            className="rotate-btn"
        >
            {label} {directionSymbol}
        </button>
    );
};

export default FaceButton;