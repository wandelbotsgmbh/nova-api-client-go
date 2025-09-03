package v2

// TrajectoryId defines model for TrajectoryId.
type TrajectoryId struct {
	// Id The identifier of the trajectory which was returned by the [addTrajectory](addTrajectory) endpoint.
	Id string `json:"id"`

	// MessageType Type specifier for server, set automatically.
	MessageType string `json:"message_type"`
}
