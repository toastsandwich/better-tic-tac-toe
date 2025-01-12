import "./UserDetails.css";
const UserDetails = ({ user }) => {
  return (
    <div className="user-details-box">
      <div className="user-detail">
        <label htmlFor="username">Username: </label>
        <span>{user.username}</span>
      </div>
      <div className="user-detail">
        <label htmlFor="country">Country: </label> <span>{user.country}</span>
      </div>
      <div className="user-detail">
        <label htmlFor="wins">Total Wins: </label>
        <span>{user.wins}</span>
      </div>
      <div className="user-detail">
        <label htmlFor="losses">Total Losses: </label>
        <span>{user.losses}</span>
      </div>
      <div className="user-detail">
        <label htmlFor="global_rank">Global Rank: </label>
        <span>{user.global_rank}</span>
      </div>
      <div className="user-detail">
        <label htmlFor="country_rank">Country Rank: </label>
        <span>{user.country_rank}</span>
      </div>
    </div>
  );
};

export default UserDetails;
