import React, { useState } from 'react'
import '../../css/top.css'

function Top() {
    const [bkimg, setBkImg] = useState('/image/BBQ.jpg')

    return (
        <div>
            <div className="Container">
                <div
                className="Container-before"
                style={{backgroundImage: `url(${bkimg})` }}
                onClick={() => setBkImg(changeBkImg(bkimg)) }
                />
            </div>
        </div>
    );
};

function changeBkImg(image) {
    if (image === '/image/BBQ.jpg') {
        return '/image/Dango.jpg'
    } else {
        return '/image/BBQ.jpg'
    }
}

export default Top;