import React from 'react';
import FaceButton from './FaceButton';

import './FaceControls.css';

const FaceControls = ({ handleRotate, loading }) => {
    const faces = [
        { face: 'front', label: 'F' },
        { face: 'right', label: 'R' },
        { face: 'up', label: 'U' },
        { face: 'back', label: 'B' },
        { face: 'left', label: 'L' },
        { face: 'down', label: 'D' }
    ];

    return (
        <>
            {faces.map(({ face, label }) => (
                <div key={face} className="face-controls">
                    <span className="face-label">{label}</span>
                    <FaceButton
                        face={face}
                        label={label}
                        direction="clockwise"
                        handleRotate={handleRotate}
                        loading={loading}
                    />
                    <FaceButton
                        face={face}
                        label={label}
                        direction="counterclockwise"
                        handleRotate={handleRotate}
                        loading={loading}
                    />
                </div>
            ))}
        </>
    );
};

export default FaceControls;