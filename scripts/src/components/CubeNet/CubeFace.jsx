import React from 'react';

const CubeFace = ({ faceData, label, position }) => {
    return (
        <div className={`face-container ${position}`}>
            <div className="face-label">{label}</div>
            <div className="face">
                {[0, 1, 2].map(row => (
                    [0, 1, 2].map(col => (
                        <div
                            key={`${row}-${col}`}
                            className={`square ${faceData[row][col]}`}
                        />
                    ))
                ))}
            </div>
        </div>
    );
};

export default CubeFace;