package ari

import "time"

// Channel represents a communication path interacting with an Asterisk server.
type Channel interface {
	// Get returns a handle to a channel for further interaction
	Get(id string) ChannelHandle

	// List lists the channels in asterisk
	List() ([]ChannelHandle, error)

	// Originate creates a new channel, returning a handle to it or an
	// error, if the creation failed
	Originate(OriginateRequest) (ChannelHandle, error)

	// Create creates a new channel, returning a handle to it or an
	// error, if the creation failed
	Create(ChannelCreateRequest) (ChannelHandle, error)

	// Data returns the channel data for a given channel
	Data(id string) (*ChannelData, error)

	// Continue tells Asterisk to return a channel to the dialplan
	Continue(id, context, extension string, priority int) error

	// Busy hangs up the channel with the "busy" cause code
	Busy(id string) error

	// Congestion hangs up the channel with the "congestion" cause code
	Congestion(id string) error

	// Answer answers the channel
	Answer(id string) error

	// Hangup hangs up the given channel
	Hangup(id string, reason string) error

	// Ring indicates ringing to the channel
	Ring(id string) error

	// StopRing stops ringing on the channel
	StopRing(id string) error

	// SendDTMF sends DTMF to the channel
	SendDTMF(id string, dtmf string, opts *DTMFOptions) error

	// Hold puts the channel on hold
	Hold(id string) error

	// StopHold retrieves the channel from hold
	StopHold(id string) error

	// Mute mutes a channel in the given direction (in,out,both)
	Mute(id string, dir Direction) error

	// Unmute unmutes a channel in the given direction (in,out,both)
	Unmute(id string, dir Direction) error

	// MOH plays music on hold
	MOH(id string, moh string) error

	// StopMOH stops music on hold
	StopMOH(id string) error

	// Silence plays silence to the channel
	Silence(id string) error

	// StopSilence stops the silence on the channel
	StopSilence(id string) error

	// Play plays the media URI to the channel
	Play(id string, playbackID string, mediaURI string) (PlaybackHandle, error)

	// Record records the channel
	Record(id string, name string, opts *RecordingOptions) (LiveRecordingHandle, error)

	// Dial dials a created channel
	Dial(id string, caller string, timeout time.Duration) error

	// Snoop spies on a specific channel, creating a new snooping channel
	Snoop(id string, snoopID string, opts *SnoopOptions) (ChannelHandle, error)

	// Subscribe subscribes on the channel events
	Subscribe(id string, n ...string) Subscription

	// Variables gets the channel Variables
	Variables(id string) Variables
}

// ChannelData is the data for a specific channel
type ChannelData struct {
	ID           string      `json:"id"`    // Unique id for this channel (same as for AMI)
	Name         string      `json:"name"`  // Name of this channel (tech/name-id format)
	State        string      `json:"state"` // State of the channel
	Accountcode  string      `json:"accountcode"`
	Caller       CallerID    `json:"caller"`    // CallerId of the calling endpoint
	Connected    CallerID    `json:"connected"` // CallerId of the connected line
	Creationtime DateTime    `json:"creationtime"`
	Dialplan     DialplanCEP `json:"dialplan"` // Current location in the dialplan
}

// ChannelCreateRequest describes how a channel should be created, when
// using the separate Create and Dial calls.
type ChannelCreateRequest struct {
	// Endpoint is the target endpoint for the dial
	Endpoint string `json:"endpoint"`

	// App is the name of the Stasis application to execute on connection
	App string `json:"app"`

	// AppArgs is the set of (comma-separated) arguments for the Stasis App
	AppArgs string `json:"appArgs,omitempty"`

	// ChannelID is the ID to give to the newly-created channel
	ChannelID string `json:"channelId,omitempty"`

	// OtherChannelID is the ID of the second created channel (when creating Local channels)
	OtherChannelID string `json:"otherChannelId,omitempty"`

	// Originator is the unique ID of the calling channel, for which this new channel-dial is being created
	Originator string `json:"originator,omitempty"`

	// Formats is the comma-separated list of valid codecs to allow for the new channel, in the case that
	// the Originator is not specified
	Formats string `json:"formats,omitempty"`
}

// SnoopOptions enumerates the non-required arguments for the snoop operation
type SnoopOptions struct {
	// App is the ARI application into which the newly-created Snoop channel should be dropped.
	App string `json:"app"`

	// AppArgs is the set of arguments to pass with the newly-created Snoop channel's entry into ARI.
	AppArgs string `json:"appArgs,omitempty"`

	// Spy describes the direction of audio on which to spy (none, in, out, both).
	// The default is 'none'.
	Spy Direction `json:"spy,omitempty"`

	// Whisper describes the direction of audio on which to send (none, in, out, both).
	// The default is 'none'.
	Whisper Direction `json:"whisper,omitempty"`
}

// ChannelHandle provides a wrapper to a Channel interface for
// operations on a particular channel ID
type ChannelHandle interface {
	// ID returns the identifier for the channel handle
	ID() string

	// Data returns the channel's data
	Data() (*ChannelData, error)

	// Continue tells Asterisk to return the channel to the dialplan
	Continue(context, extension string, priority int) error

	// Play initiates playback of the specified media uri
	// to the channel, returning the Playback handle
	Play(id string, mediaURI string) (PlaybackHandle, error)

	// Record records the channel to the given filename
	Record(name string, opts *RecordingOptions) (LiveRecordingHandle, error)

	// Busy hangs up the channel with the "busy" cause code
	Busy() error

	// Congestion hangs up the channel with the congestion cause code
	Congestion() error

	// Hangup hangs up the channel with the normal cause code
	Hangup() error

	// Answer answers the channel
	Answer() error

	// IsAnswered checks the current state of the channel to see if it is "Up"
	IsAnswered() (bool, error)

	// Ring indicates ringing to the channel
	Ring() error

	// StopRing stops ringing on the channel
	StopRing() error

	// Mute mutes the channel in the given direction (in, out, both)
	Mute(dir Direction) (err error)

	// Unmute unmutes the channel in the given direction (in, out, both)
	Unmute(dir Direction) (err error)

	// Hold puts the channel on hold
	Hold() error

	// StopHold retrieves the channel from hold
	StopHold() error

	// MOH plays music on hold of the given class
	// to the channel
	MOH(mohClass string) error

	// StopMOH stops playing of music on hold to the channel
	StopMOH() error

	// Variables returns the channel variables
	Variables() Variables

	// Originate creates (and dials) a new channel using the present channel as its Originator.
	Originate(req OriginateRequest) (ChannelHandle, error)

	// Dial dials a created channel.  `caller` is the optional
	// channel ID of the calling party (if there is one).  Timeout
	// is the length of time to wait before the dial is answered
	// before aborting.
	Dial(caller string, timeout time.Duration) error

	// Snoop spies on a specific channel, creating a new snooping channel placed into the given app
	Snoop(snoopID string, opts *SnoopOptions) (ChannelHandle, error)

	// Silence plays silence to the channel
	Silence() error

	// StopSilence stops silence to the channel
	StopSilence() error

	// Subscribe subscribes the list of channel events
	Subscribe(n ...string) Subscription

	// SendDTMF sends the DTMF information to the server
	SendDTMF(dtmf string, opts *DTMFOptions) error

	// Match returns true if the event matches the channel
	Match(e Event) bool
}
