//This file is a react component, uses typescript
//Create component here. Ammaad Denmark 11/13
import React from 'react';
import styles from "./styles.module.css";
import Link from 'next/link';
import PropTypes from 'prop-types';

interface IMajorListItem {
    majorName: string,
    description: string,
    path:string//Learn more button
}

const MajorListItem: React.FC<IMajorListItem> = ({majorName, description, path}) => {

    return (//defines what will be displayed.
        <div className={styles.majorsListItem}>
            <h2>{majorName}</h2>
            <p>Description:{description}</p>
            <Link href={path}>Learn more</Link>
        </div>
    )

}

MajorListItem.PropTypes= {//CHECK
    majorName: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
    path: PropTypes.string.isRequired,
}
//End of react component, add class name, styles with css
//Look at other components for reference. Ask about path for learn more button


export default MajorListItem;//Import into other parts of project