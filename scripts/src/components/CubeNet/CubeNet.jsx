import React from 'react';
import CubeFace from './CubeFace';
import './CubeNet.css';

const CubeNet = ({ cubeState }) => {
    const faceConfig = [
        { face: 'up', label: 'U', position: 'up' },
        { face: 'left', label: 'L', position: 'left' },
        { face: 'front', label: 'F', position: 'front' },
        { face: 'right', label: 'R', position: 'right' },
        { face: 'back', label: 'B', position: 'back' },
        { face: 'down', label: 'D', position: 'down' }
    ];

    return (
        <div className="cube-net">
            {faceConfig.map(config => (
                <CubeFace
                    key={config.face}
                    face={config.face}
                    faceData={cubeState[config.face]}
                    label={config.label}
                    position={config.position}
                />
            ))}
        </div>
    );
};

export default CubeNet;