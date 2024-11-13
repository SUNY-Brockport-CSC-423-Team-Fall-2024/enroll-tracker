"use client";

import { headerLinks, footerLinks } from "../lib/data";
import styles from "./styles.module.css";
import Link from "next/link";


interface pageSliderData {
    firstPage: string;
    secondPage: string;
} // end of pageSliderData

const PageSlider: React.FC<pageSliderData> = ({ firstPage, secondPage }) => {
    return (
        <div className={styles.full_slider}>
        <div className={styles.slider_left}>
            <div className={styles.slider_content}>
                <h3>{firstPage}</h3>
            </div>
        </div>
        <div className={styles.slider_right}>
            <div className={styles.slider_content}>
                <h3>{secondPage}</h3>
            </div>
        </div>
        </div>); // end of return pageSlider
}; // end of pageSlider

export default PageSlider;