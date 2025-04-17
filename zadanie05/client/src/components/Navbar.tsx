import {Link} from "react-router-dom";

type Props = {
    links: string[],
    labels: string[]
}

const Navbar = ({links, labels} : Props) => {

    return (
    <nav className="bar">
        <ul style={{ display: 'flex', gap: '1rem', listStyle: 'none' }}>
            {links.map((link, i) => (
                <li key={link}>
                    <Link to={link}>{labels[i]}</Link>
                </li>
            ))}
        </ul>
    </nav>
    )

}

export default Navbar;