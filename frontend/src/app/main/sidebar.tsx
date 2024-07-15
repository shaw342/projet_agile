// components/Sidebar.js
import styles from './Sidebar.module.css';
import { IoChatbubbleEllipsesOutline } from "react-icons/io5";

const Sidebar = () => {
    return (
        <div className={styles.sidebar}>
            <div className={styles.sidebarBrand}>
                <a href="#">MaMarque</a>
            </div>
            <ul className={styles.sidebarNav}>
                <li><a href="#home">Accueil</a></li>
                <li><a href="#services">notes</a></li>
                <li><a href="#about">calandrier</a></li>
                <li><a href="#contact"><IoChatbubbleEllipsesOutline /> Chat</a></li>
            </ul>
        </div>
    );
};

export default Sidebar;
