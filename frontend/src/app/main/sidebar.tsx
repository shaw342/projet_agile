
import styles from './Sidebar.module.css';

function Sidebar() {
  return (
    <nav className={styles.sidebar}>
      <ul>
        <li className={styles.sidebarItem}>
            <a>Home</a>
        </li>
        <li className={styles.sidebarItem}>
            <a>About</a>
          
        </li>
        <li className={styles.sidebarItem}>
            <a>Contact</a>
        </li>
      </ul>
    </nav>
  );
}

export default Sidebar;

