//This is the component for courses listed under major, right??
//Ammaad Denmark 11/13
import React from 'react';
import PropTypes from 'prop-types';
import styles from "./styles.module.css";

interface ICourseListItem {
    courseName: string,
    description: string,
}

const CourseListItem: React.FC<ICourseListItem> = ({courseName, description}) => {

    return (//defines what will be displayed.Class name?
        <div>
            <h2>{courseName}</h2>
            <p>Description:{description}</p>
        </div>
    )

}

CourseListItem.PropTypes= {//CHECK
    courseName: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
}